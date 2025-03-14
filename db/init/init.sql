
CREATE DATABASE public;

CREATE TABLE Users (
	UserId SERIAL primary key,
	username text,
	password text
);

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
    PartnerId TEXT,
    FOREIGN KEY (PartnerId) REFERENCES Partners(PartnerId)
);

CREATE TABLE Skus (
    SkuId TEXT PRIMARY KEY,
    SkuName TEXT,
    AvailabilityId TEXT
);

CREATE TABLE Products (
    ProductId TEXT PRIMARY KEY,
    SkuId TEXT,
    ProductName TEXT,    
    PublisherId TEXT,
    PublisherName TEXT,
    SubscriptionId TEXT,
    SubscriptionDescription text,
    
    FOREIGN KEY (SkuId) REFERENCES Skus(SkuId)
    
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
);

CREATE TABLE Billings (
    BillingId SERIAL PRIMARY KEY, 
    CustomerId TEXT,
    ProductId text,
    EntitlementId TEXT,
    InvoiceNumber TEXT,
    ChargeStartDate DATE,
    ChargeEndDate DATE,
    UsageDate DATE,
    MeterId TEXT,
    MeterType TEXT,
    MeterCategory TEXT,
    MeterSubCategory TEXT,
    MeterName TEXT,
    MeterRegion TEXT,
    Unit TEXT,
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
    AdditionalInfo JSON,
    EffectiveUnitPrice DECIMAL(18, 6),
    PCToBCExchangeRate DECIMAL(18, 6),
    PCToBCExchangeRateDate DATE,
    FOREIGN KEY (CustomerId) REFERENCES Customers(CustomerId),
    FOREIGN KEY (ProductId) REFERENCES Products(ProductId),
    FOREIGN KEY (EntitlementId) REFERENCES Entitlements(EntitlementId)
);