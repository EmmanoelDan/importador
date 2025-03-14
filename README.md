# Importador API 🗂️

## 📌 Sobre

A **Importador API** permite o upload e processamento de arquivos, além de autenticação de usuários via **JWT**. Seu objetivo é facilitar a importação e gerenciamento de dados com performance, eficiência e segurança.

## 🚀 Tecnologias

- **Golang** (Gin, GORM)
- **PostgreSQL/MySQL**
- **Docker**
- **JWT** para autenticação
- **GODOTENV** para variaveis de ambiente

## 📜 Instalação

1️⃣ Clone o repositório:

```sh
git clone https://github.com/seu-usuario/importador.git
cd importador
```

2️⃣ Instale as dependências:

```sh
go mod tidy
```

3️⃣ Configure o arquivo .env com as variáveis de ambiente:

```sh
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
JWT_KEY=
```

5️⃣ Execute o docker-compose.yml:

obs: certifique-se de que o docker esteja funcionando corretamente em sua maquina

```sh
docker-compose up --build -d
```

```sh
    http://localhost:8080
```

## 📌 Endpoints

OBS: POST Register -> POST Sign -> POST Import Files -> GET Bellings

- **URL:** `/register`
- **Método:** `POST`

### Body:

<pre>
{
  "username": "admin",
  "password": "123456"
}
</pre>

### Resposta (201 created):

<pre>
{
    "message": "Create user successfully",
    "user": "admin"
}
</pre>

- **URL:** `/login`
- **Método:** `POST`

### Body:

<pre>

{
  "username": "admin",
  "password": "123456"
}
</pre>

### Resposta (200 OK):

<pre>
{
  token": "jwt_token_aqui"
}
</pre>

- **URL:** `/import_file`
- **Método:** `POST`
- **Tipo de Conteúdo:** `multipart/form-data`
- **Autenticação:** Bearer Token (JWT)
<pre>
{
    "message": "Upload successfully imported"
}
</pre>

- **URL:** `/billings?page=1&pageSize=10`
- **Método:** `GET`
- **Tipo de Conteúdo:** `multipart/form-data`
- **Autenticação:** Bearer Token (JWT)
- **Params** pagination: page e pageSize.

<pre>
{
    "billings": [
        {
            "BillingId": "1",
            "CustomerId": "12345678-abcd-4321-abcd-1234abcdef1234",
            "ProductId": "XYZ1234567",
            "EntitlementId": "98765abcd-1234-5678-9876-abcdef123456",
            "InvoiceNumber": "INV123456789",
            "ChargeStartDate": "2023-01-01T00:00:00Z",
            "ChargeEndDate": "2023-01-02T00:00:00Z",
            "UsageDate": "2023-01-01T00:00:00Z",
            "MeterId": "abcd1234-5678-9876-abc1-2345xyz9876",
            "MeterType": "10GB Data Transfer",
            "MeterCategory": "Cloud Storage",
            "MeterSubCategory": "Data Ingestion",
            "MeterName": "Data Transfer Basic Plan",
            "MeterRegion": "USWest",
            "Unit": "GB",
            "ResourceLocation": "USWest",
            "ConsumedService": "Cloud.Storage",
            "ResourceGroup": "cloudservices",
            "ResourceURI": "/subscriptions/12345678-abcd-1234-abcd-1234abcdef1234/resourceGroups/cloudservices/providers/Cloud.Storage/dataTransfer/standardPlan",
            "ChargeType": "Recurring",
            "UnitPrice": 0.005,
            "Quantity": 50,
            "UnitType": "GB",
            "BillingPreTaxTotal": 0.25,
            "BillingCurrency": "USD",
            "PricingPreTaxTotal": 0.25,
            "PricingCurrency": "USD",
            "ServiceInfo1": "Standard Plan",
            "ServiceInfo2": "Monthly Billing",
            "Tags": null,
            "AdditionalInfo": {
                "IpAddress": "192.168.0.1",
                "PlatformType": "CloudV2"
            },
            "EffectiveUnitPrice": 0.0045,
            "PCToBCExchangeRate": 1.0,
            "PCToBCExchangeRateDate": "2023-01-01T00:00:00Z",
            "Customer": {
                "CustomerId": "12345678-abcd-4321-abcd-1234abcdef1234",
                "CustomerName": "TechCorp",
                "CustomerDomainName": "techcorp.com",
                "CustomerCountry": "US",
                "PartnerId": "abcdef12-3456-7890-abcd-1234abcdef5678",
                "Partner": {
                    "PartnerId": "abcdef12-3456-7890-abcd-1234abcdef5678",
                    "PartnerName": "TechSolutions LLC",
                    "MpnId": "54321",
                    "Tier2MpnId": "987654"
                }
            },
            "Product": {
                "ProductId": "XYZ1234567",
                "SkuId": "100",
                "ProductName": "Cloud Data Transfer",
                "PublisherName": "TechCloud",
                "PublisherId": "tccloudpub",
                "SubscriptionId": "a1b2c3d4-e5f6-7890-abcd-1234567890ab",
                "SubscriptionDescription": "Standard Cloud Subscription",
                "Sku": {
                    "SkuId": "100",
                    "SkuName": "Premium",
                    "AvailabilityId": "XYZ987654321"
                }
            },
            "Entitlement": {
                "EntitlementId": "98765abcd-1234-5678-9876-abcdef123456",
                "EntitlementDescription": "Cloud Storage Premium",
                "PartnerEarnedCreditPercentage": 10,
                "CreditPercentage": 5,
                "CreditType": "Customer Credit",
                "BenefitOrderId": "1001",
                "BenefitId": "A1B2C3D4",
                "BenefitType": "Discount"
            }
        }
    ]
}

</pre>

📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.
