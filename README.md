### 💡 This project was built by [cape4Labs](https://cape4.tech/) specifically for the [BUIDL BATTLE #2](https://dorahacks.io/hackathon/buidlbattle2) hackathon.

---

## 🌁 Capebridge

**Capebridge** is a gateway to swap **Base USDC tokens** for **Stacks USDCX tokens** (and vice versa). 

It is a decentralized bridge service controlled by **a federation of nodes**.

The **testnet** version is **already deployed** - you can try it here: capebridge.tech. 

## 🌉 How It Works

**The protocol works as follows:**

- A smart contract is deployed on each supported blockchain, allowing users to deposit their tokens. Once a deposit is made, the tokens are locked and a corresponding bridge event is emitted.

- Federation nodes are listening for these events. Each node independently verifies the deposit and relays the data to the destination chain.

- On the destination blockchain, the smart contract aggregates and compares the submitted data. Once 2 out of 3 nodes reach quorum on a given bridge request, the transfer is approved and a request is executed to release tokens to the user.

- **There are no mint or burn mechanics involved, and no wrapped assets are used.** Users always receive **official Circle tokens** — either the widely used USDC on Base or the newly introduced USDCX on Stacks.

## 🌃 Architecture 

Below you can find how the bridge works in both directions step-by-step:

<img width="1855" height="562" alt="image_2026-03-20_20-52-55" src="https://github.com/user-attachments/assets/138d6e71-d5b7-4953-9ef4-7f5370ebd34f" />

<img width="1852" height="566" alt="image_2026-03-20_20-52-55 (2)" src="https://github.com/user-attachments/assets/59a03d19-fa09-4e06-a0d8-2d15beacb52c" />

**Thus, our project consists of:**

**(x3) Node:**
 - Stacks/Base indexer and relayer
 - Postgres db
 - Stacks signer 
 
**(x1) DApp:**
 - Golang SSE backend
 - Nextjs frontend
 
**(x1) Smart Contracts:**
 - Stacks bridge + federation contract
 - Base federation contract
 - Base bridge contract

## 🌌 Production Release

**This project is a demonstration of bridging capabilities between Stacks and Base, built specifically for the hackathon.** 
It is not intended to be used as a production-ready solution, and should not be considered a template for deploying real bridges on mainnet.

- Extend admin functionality by adding a broader set of controls for protocol management, monitoring, and emergency actions

- Implement audits both for Clarity and Solidity smart contracts.

- Decentralize node operations by onboarding trusted independent operators.

- Introduce operator fees and define a clear fee model (recommended ~0.4% per transaction).

- Define and enforce minimum and maximum bridge limits to prevent misuse and edge-case risks.

**Only after completing all the steps above can the bridge be considered reasonably safe for mainnet deployment and handling user funds.**

---

### __If you loved this project, star it ⭐️ and check other cape4labs repos [here](https://github.com/cape4labs)!__
