# Features
CASHBACK = "cashback"
MIXED_PAYMENTS = "mixed payments"
PAYMENT = "payments"
REFUND = "refunds"
SETTINGS = "settings"
WALLET = "wallet"
ALLTOPUP = "all top ups"
INLINETOPUP = "inline topup"
MANUALTOPUP = "manual topup"
AUTOTOPUP = "auto topup"
SECURITY = "wallet_security"

# keyword: feature
features: dict[str, str] = {
    "AUTO_TOP_UP": AUTOTOPUP,
    "INLINE_TOP_UP": INLINETOPUP,
    "COD_NO_BALANCE": MIXED_PAYMENTS,
    "COD_BALANCE": MIXED_PAYMENTS,
    "CVC": PAYMENT,
    "MM_YY": PAYMENT,
    "SETTING": SETTINGS,
    "FAST_TOPUP": INLINETOPUP,
    "LIMIT": WALLET,
    "TOPUP": ALLTOPUP,
    "CASHBACK": CASHBACK,
    "XENDIT": PAYMENT,
    "_PIN_": SECURITY,
    "PAYMENT": PAYMENT,
    "REFUND": REFUND,
    "TRANSFER": REFUND,
    "TRANSACTION_HISTORY": WALLET,
    "CARD": PAYMENT,
    "BALANCE": WALLET,
    "REMAINING_AMOUNT": WALLET,
    "PANDAPAY": WALLET,
    "SAVE": WALLET,
    "AVAILABLE_CREDIT": PAYMENT,
    "WALLET": WALLET,
    "WALLET_ACTIVITY": WALLET,
}

featureFreq: dict[str, int] = {
    CASHBACK: 0,
    MIXED_PAYMENTS: 0,
    PAYMENT: 0,
    REFUND: 0,
    SETTINGS: 0,
    WALLET: 0,
    ALLTOPUP: 0,
    INLINETOPUP: 0,
    MANUALTOPUP: 0,
    AUTOTOPUP: 0,
    SECURITY: 0
}

with open("./translation_keys") as f:
    for line in f:
        found = False
        line = line.upper()
        line = line.strip()
        for keyword in features:
            if keyword in line:
                print(f"{line}, {features[keyword]}")
                found = True
                featureFreq[features[keyword]] += 1
                break  # only 1 match per translation key
        if not found:
            print(f"{line}: no feature")
print("=" * 50)
for k, v in featureFreq.items():
    print(f"{k} feature: {v}")
