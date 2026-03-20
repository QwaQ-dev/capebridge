(use-trait sip-010-trait .sip-010-trait.sip-010-trait)


(define-constant federation-node-1 'ST30CQ2DH1ZN7T36S338SFTZM92Q6JZYDTKEFW4TB)
(define-constant federation-node-2 'STNBDYJPTWBE0NRZWKZKM2BRCWRRG0W2T252YPCJ)
(define-constant federation-node-3 'ST2YYVR9BA931YQKTCHHV06CW9C8XNG31TVN93JQC)
(define-constant contract-address  'ST1PGPX7HRMFBBXFKG6YCNDCKN5TZJF32VMF443ZQ.capebridge-test-v1)
(define-constant admin-address  'ST6BVG3ADY71WYYTZKFNYXCE5604JX3KG7N7R17X)




(define-constant ERR-INVALID-AMOUNT         (err u100))
(define-constant ERR-INSUFFICIENT-EXTERNAL  (err u101))
(define-constant ERR-INSUFFICIENT-OWN       (err u104))
(define-constant ERR-ALREADY-TRANSFERRED    (err u200))
(define-constant ERR-ALREADY-CONFIRMED-1    (err u201))
(define-constant ERR-ALREADY-CONFIRMED-2    (err u202))
(define-constant ERR-ALREADY-CONFIRMED-3    (err u203))
(define-constant ERR-NOT-FEDERATION-NODE    (err u204))
(define-constant ERR-DATA-MISMATCH          (err u205))
(define-constant ERR-NO-CONSENSUS           (err u206))
(define-constant ERR-REQUEST-NOT-FOUND      (err u207))
(define-constant ERR-ALREADY-INITIALIZED    (err u300))


(define-data-var is-initialized   bool false)
(define-data-var own-balance      uint u0)
(define-data-var external-balance uint u0)
(define-data-var start-balance    uint u0)
(define-data-var nonce            uint u0)


(define-map PendingRequests
  uint
  {
    transfer-made:        bool,
    node-1-confirmation:  bool,
    node-2-confirmation:  bool,
    node-3-confirmation:  bool,
    node-1-recipient:     principal,
    node-2-recipient:     principal,
    node-3-recipient:     principal,
    node-1-amount:        uint,
    node-2-amount:        uint,
    node-3-amount:        uint
  }
)


(define-private (empty-request)
  {
    transfer-made:        false,
    node-1-confirmation:  false,
    node-2-confirmation:  false,
    node-3-confirmation:  false,
    node-1-recipient:     tx-sender,
    node-2-recipient:     tx-sender,
    node-3-recipient:     tx-sender,
    node-1-amount:        u0,
    node-2-amount:        u0,
    node-3-amount:        u0
  }
)


(define-private (is-federation-node)
  (or
    (is-eq tx-sender federation-node-1)
    (is-eq tx-sender federation-node-2)
    (is-eq tx-sender federation-node-3)
  )
)

(define-private (is-admin-node)
  (or
    (is-eq tx-sender federation-node-1)
    (is-eq tx-sender federation-node-2)
    (is-eq tx-sender federation-node-3)
    (is-eq tx-sender admin-address)
  )
)



(define-public (init (own uint) (external uint))
  (begin
    (asserts! (not (var-get is-initialized)) ERR-ALREADY-INITIALIZED)
    (asserts! (is-admin-node) ERR-NOT-FEDERATION-NODE)
    (var-set own-balance own)
    (var-set external-balance external)
    (var-set start-balance (+ own external))
    (var-set is-initialized true)
    (ok true)
  )
)


(define-private (bridge-request (amount uint) (receiver (string-ascii 100)) (sender principal))
  (begin
    (asserts! (> (var-get external-balance) amount) ERR-INSUFFICIENT-EXTERNAL)
    (var-set external-balance (- (var-get external-balance) amount))
    (var-set own-balance (+ (var-get own-balance) amount))
    (print {
      event:    "request-approved",
      sender:   sender,
      amount:   amount,
      receiver: receiver,
      nonce:    (var-get nonce)
    })
    (var-set nonce (+ (var-get nonce) u1))
    (ok true)
  )
)


(define-public (deposit (amount uint) (receiver (string-ascii 100)) (token <sip-010-trait>))
  (begin
    (asserts! (> amount u0) ERR-INVALID-AMOUNT)
    (try! (contract-call? token transfer amount tx-sender contract-address none))
    (try! (bridge-request amount receiver tx-sender))
    (ok true)
  )
)


(define-private (check-data-consensus
    (req { transfer-made: bool,
           node-1-confirmation: bool, node-2-confirmation: bool, node-3-confirmation: bool,
           node-1-recipient: principal, node-2-recipient: principal, node-3-recipient: principal,
           node-1-amount: uint, node-2-amount: uint, node-3-amount: uint })
    (recipient principal)
    (amount uint))
  (and
    (if (get node-1-confirmation req)
      (and (is-eq (get node-1-recipient req) recipient)
           (is-eq (get node-1-amount req) amount))
      true)
    (if (get node-2-confirmation req)
      (and (is-eq (get node-2-recipient req) recipient)
           (is-eq (get node-2-amount req) amount))
      true)
    (if (get node-3-confirmation req)
      (and (is-eq (get node-3-recipient req) recipient)
           (is-eq (get node-3-amount req) amount))
      true)
  )
)


(define-public (confirm-request (recipient principal) (amount uint) (request-nonce uint))
  (let (
    (req (default-to (empty-request) (map-get? PendingRequests request-nonce)))
  )
    (asserts! (not (get transfer-made req)) ERR-ALREADY-TRANSFERRED)

    (if (is-eq tx-sender federation-node-1)
      (begin
        (asserts! (not (get node-1-confirmation req)) ERR-ALREADY-CONFIRMED-1)
        (map-set PendingRequests request-nonce (merge req {
          node-1-confirmation: true,
          node-1-recipient:    recipient,
          node-1-amount:       amount
        }))
      )
      (if (is-eq tx-sender federation-node-2)
        (begin
          (asserts! (not (get node-2-confirmation req)) ERR-ALREADY-CONFIRMED-2)
          (map-set PendingRequests request-nonce (merge req {
            node-2-confirmation: true,
            node-2-recipient:    recipient,
            node-2-amount:       amount
          }))
        )
        (if (is-eq tx-sender federation-node-3)
          (begin
            (asserts! (not (get node-3-confirmation req)) ERR-ALREADY-CONFIRMED-3)
            (map-set PendingRequests request-nonce (merge req {
              node-3-confirmation: true,
              node-3-recipient:    recipient,
              node-3-amount:       amount
            }))
          )
          (asserts! false ERR-NOT-FEDERATION-NODE)
        )
      )
    )

    (let (
      (updated-req (default-to (empty-request) (map-get? PendingRequests request-nonce)))
    )
      (asserts! (check-data-consensus updated-req recipient amount) ERR-DATA-MISMATCH)
    )

    (print {
      event:     "request-confirmed",
      sender:    tx-sender,
      recipient: recipient,
      amount:    amount,
      nonce:     request-nonce
    })

    (ok true)
  )
)


(define-private (has-consensus (request-nonce uint))
  (let (
    (req (default-to (empty-request) (map-get? PendingRequests request-nonce)))
    (c1 (get node-1-confirmation req))
    (c2 (get node-2-confirmation req))
    (c3 (get node-3-confirmation req))
    (r1 (get node-1-recipient req))
    (r2 (get node-2-recipient req))
    (r3 (get node-3-recipient req))
    (a1 (get node-1-amount req))
    (a2 (get node-2-amount req))
    (a3 (get node-3-amount req))
  )
    (or
      (and c1 c2 (is-eq r1 r2) (is-eq a1 a2))
      (and c1 c3 (is-eq r1 r3) (is-eq a1 a3))
      (and c2 c3 (is-eq r2 r3) (is-eq a2 a3))
    )
  )
)


(define-private (get-consensus-recipient (request-nonce uint))
  (let (
    (req (default-to (empty-request) (map-get? PendingRequests request-nonce)))
  )
    (if (and (get node-1-confirmation req) (get node-2-confirmation req)
             (is-eq (get node-1-recipient req) (get node-2-recipient req)))
      (get node-1-recipient req)
      (if (and (get node-1-confirmation req) (get node-3-confirmation req)
               (is-eq (get node-1-recipient req) (get node-3-recipient req)))
        (get node-1-recipient req)
        (get node-2-recipient req)
      )
    )
  )
)


(define-private (get-consensus-amount (request-nonce uint))
  (let (
    (req (default-to (empty-request) (map-get? PendingRequests request-nonce)))
  )
    (if (and (get node-1-confirmation req) (get node-2-confirmation req)
             (is-eq (get node-1-amount req) (get node-2-amount req)))
      (get node-1-amount req)
      (if (and (get node-1-confirmation req) (get node-3-confirmation req)
               (is-eq (get node-1-amount req) (get node-3-amount req)))
        (get node-1-amount req)
        (get node-2-amount req)
      )
    )
  )
)


(define-public (transfer (request-nonce uint) (token <sip-010-trait>))
  (let (
    (req       (unwrap! (map-get? PendingRequests request-nonce) ERR-REQUEST-NOT-FOUND))
    (recipient (get-consensus-recipient request-nonce))
    (amount    (get-consensus-amount request-nonce))
  )
    (asserts! (is-federation-node) ERR-NOT-FEDERATION-NODE)
    (asserts! (not (get transfer-made req)) ERR-ALREADY-TRANSFERRED)
    (asserts! (has-consensus request-nonce) ERR-NO-CONSENSUS)
    (asserts! (<= amount (var-get own-balance)) ERR-INSUFFICIENT-OWN)

    (try! (contract-call? token transfer amount contract-address recipient none))

    (var-set own-balance (- (var-get own-balance) amount))
    (map-set PendingRequests request-nonce (merge req { transfer-made: true }))

    (print {
      event:     "transfer-executed",
      recipient: recipient,
      amount:    amount,
      nonce:     request-nonce
    })

    (ok true)
  )
)


(define-read-only (get-request (request-nonce uint))
  (map-get? PendingRequests request-nonce)
)


(define-read-only (get-own-balance)      (var-get own-balance))
(define-read-only (get-external-balance) (var-get external-balance))
(define-read-only (get-nonce)            (var-get nonce))
