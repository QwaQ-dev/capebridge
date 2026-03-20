const { 
    makeContractCall, 
    broadcastTransaction, 
    AnchorMode, 
    PostConditionMode, 
    uintCV, 
    principalCV,
    contractPrincipalCV 
} = require('@stacks/transactions');

// Import network configuration (Stacks Testnet)
const { STACKS_TESTNET } = require('@stacks/network'); 

const express = require('express');

const app = express();
app.use(express.json());

const PORT = process.env.PORT || 3000;
const PRIVATE_KEY = process.env.STACKS_PRIVATE_KEY;

// HTTP endpoint to sign and broadcast a contract call
app.post('/sign-contract-call', async (req, res) => {
    const { contractAddress, contractName, functionName, functionArgs } = req.body;

    try {
        // Map incoming generic args into Clarity Values (CV)
        const mappedArgs = functionArgs.map(arg => {

            // Unsigned integer - Clarity uint
            if (arg.type === 'uint') return uintCV(arg.value);

            if (arg.type === 'principal') {
                // If string contains ".", treat it as contract principal (address.contract-name)
                if (arg.value.includes('.')) {
                    const [addr, name] = arg.value.split('.');
                    return contractPrincipalCV(addr, name);
                }

                // Otherwise - standard account principal
                return principalCV(arg.value);
            }

            // Explicit trait handling (same as contract principal in practice)
            if (arg.type === 'trait') {
                const [addr, name] = arg.value.split('.');
                return contractPrincipalCV(addr, name);
            }

            // Fail fast on unsupported types (important for safety)
            throw new Error(`Unsupported arg type: ${arg.type}`);
        });

        // Build transaction options
        const txOptions = {
            contractAddress,
            contractName,
            functionName,
            functionArgs: mappedArgs,

            // Private key used to sign transaction
            senderKey: PRIVATE_KEY,

            // IMPORTANT:
            // Disabled ABI validation - you rely entirely on correct arg encoding
            validateWithAbi: false,

            network: STACKS_TESTNET,

            // AnchorMode.Any → allows inclusion in microblocks or anchor blocks
            anchorMode: AnchorMode.Any,

            // No post-conditions enforced (unsafe but flexible)
            postConditionMode: PostConditionMode.Allow,
        };

        // Create signed transaction
        const transaction = await makeContractCall(txOptions);

        // Broadcast transaction to Stacks node
        const result = await broadcastTransaction({
            transaction: transaction,
            network: STACKS_TESTNET
        });

        console.log(`TX [${functionName}] broadcasted. ID: ${result.txid}`);
        
        return res.json(result);

    } catch (error) {
        console.error("Signer error:", error);

        // Prevent double response (Express safeguard)
        if (!res.headersSent) {
            res.status(500).json({ error: error.message });
        }
    }
});

// Bind server to all interfaces (required for Docker / remote access)
app.listen(PORT, '0.0.0.0', () => {
    console.log(`Stacks Signer running on port ${PORT}`);
});