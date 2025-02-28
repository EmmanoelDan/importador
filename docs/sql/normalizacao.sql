
CREATE TABLE Partners (
    PartnerId TEXT PRIMARY KEY,
    PartnerName TEXT,
    MpnId TEXT,
    Tier2MpnId TEXT
);

CREATE TABLE Customers (
    CustomerId TEXT PRIMARY KEY,
    CustomerName TEXT,
    CustomerDomainName TEXT,
    CustomerCountry TEXT,
    PartnerId VARCHAR(255),
    CONSTRAINT fk_partner FOREIGN KEY (PartnerId) REFERENCES Partners(PartnerId)
);

CREATE TABLE Products (
    ProductId TEXT PRIMARY KEY,
    ProductName TEXT,    
    PublisherId TEXT,
    PublisherName TEXT
);

CREATE TABLE Skus (
	SkuId TEXT PRIMARY KEY,
	ProductId TEXT,
	SkuName TEXT,
	AvailabilityId TEXT,
	FOREIGN KEY (ProductId) REFERENCES Products(ProductId)
)

CREATE TABLE Subscriptions (
    SubscriptionId TEXT PRIMARY KEY,
    CustomerId TEXT,
    ProductId TEXT,
    SubscriptionDescription TEXT,
    FOREIGN KEY (CustomerId) REFERENCES Customers(CustomerId),
    FOREIGN KEY (ProductId) REFERENCES Products(ProductId)
);

CREATE TABLE Meters (
    MeterId TEXT PRIMARY KEY,
    MeterType TEXT,
    MeterCategory TEXT,
    MeterSubCategory TEXT,
    MeterName TEXT,
    MeterRegion TEXT,
    Unit TEXT
);

CREATE TABLE Billings (
    BillingId TEXT PRIMARY KEY AUTO_INCREMENT,
    PartnerId TEXT,
    SubscriptionId TEXT,
    MeterId TEXT,
    EntitlementId TEXT,
    InvoiceNumber TEXT,
    ChargeStartDate DATE,
    ChargeEndDate DATE,
    UsageDate DATE,
    ResourceLocation TEXT,
    ConsumedService TEXT,
    ResourceGroup TEXT,
    ResourceURI TEXT,
    ChargeType TEXT,
    UnitPrice DECIMAL(18, 6),
    Quantity DECIMAL(18, 6),
    UnitType TEXT,
    BillingPreTaxTotal DECIMAL(18, 6),
    BillingCurrency TEXT,
    PricingPreTaxTotal DECIMAL(18, 6),
    PricingCurrency TEXT,
    ServiceInfo1 TEXT,
    ServiceInfo2 TEXT,
    Tags JSON,
    AdditionalInfo JSOB,
    EffectiveUnitPrice DECIMAL(18, 6),
    PCToBCExchangeRate DECIMAL(18, 6),
    PCToBCExchangeRateDate DATE,
    FOREIGN KEY (PartnerId) REFERENCES Partners(PartnerId),
    FOREIGN KEY (SubscriptionId) REFERENCES Subscriptions(SubscriptionId),
    FOREIGN KEY (MeterId) REFERENCES Meters(MeterId),
    FOREIGN KEY (EntitlementId) REFERENCES Entitlements(EntitlementId),
);


CREATE TABLE Entitlements (
    EntitlementId TEXT PRIMARY KEY,
    EntitlementDescription TEXT,
    PartnerEarnedCreditPercentage DECIMAL(5, 2),
    CreditPercentage DECIMAL(5, 2),
    CreditType TEXT,
    BenefitOrderId TEXT,
    BenefitId TEXT,
    BenefitType TEXT

) 


