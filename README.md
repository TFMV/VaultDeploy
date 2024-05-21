# VaultDeploy

VaultDeploy is a user-friendly API designed to deploy Data Vault 2.0 models into PostgreSQL from CSV files, simplifying schema management and data integration.

![FuzzyMatchFinder](assets/vaultdeploy.webp)

## Introduction

Data Vault 2.0 is a methodology and architecture for data warehousing that addresses the challenges of agility, scalability, and adaptability in data integration. VaultDeploy leverages this methodology to help users seamlessly deploy and manage their Data Vault models in a PostgreSQL database.

## Key Features

- **Automated Schema Management:** Automatically create or update Data Vault 2.0 structures (hubs, links, satellites) based on the provided CSV file.
- **User-Friendly Configuration:** Simplified configuration through a YAML file, making it easy for users to specify database connection details and Data Vault structures.
- **Web Interface:** Interactive web interface for uploading CSV files and configuring Data Vault models without manual file edits.
- **Error Handling:** Comprehensive error handling and user feedback to ensure smooth data integration processes.

## How Data Vault 2.0 Works

### Components of Data Vault 2.0

1. **Hubs:** 
   - Represent core business entities.
   - Contain unique business keys with minimal descriptive information.
   - Ensure the integrity and uniqueness of business keys.

2. **Links:**
   - Define relationships between hubs.
   - Capture the many-to-many relationships between business entities.
   - Allow for the association of multiple business keys from different hubs.

3. **Satellites:**
   - Store descriptive information and contextual data for hubs and links.
   - Contain metadata such as timestamps and record sources.
   - Enable historical tracking and auditing of changes over time.

### Advantages of Data Vault 2.0

- **Scalability:** Designed to handle large volumes of data and support parallel loading processes.
- **Flexibility:** Easily adapt to changing business requirements and integrate new data sources without major schema changes.
- **Auditability:** Maintain historical records and metadata to ensure data integrity and traceability.
- **Agility:** Support rapid development and deployment cycles, enabling quick response to business needs.

## Use Cases

- **Enterprise Data Warehousing:** Build scalable and flexible data warehouses that can adapt to evolving business requirements.
- **Data Integration:** Seamlessly integrate data from multiple sources while maintaining historical accuracy and auditability.
- **Business Intelligence:** Provide a solid foundation for business intelligence and analytics by ensuring data consistency and reliability.
- **Regulatory Compliance:** Ensure compliance with data governance and regulatory requirements through robust auditing and historical tracking.

## Getting Started

1. **Clone the Repository:**
   - Download the VaultDeploy repository from GitHub and navigate to the project directory.

2. **Update Configuration:**
   - Edit the `config.yaml` file to specify your PostgreSQL connection details and Data Vault structures.

3. **Run the Application:**
   - Start the VaultDeploy application and access the web interface for uploading CSV files and configuring Data Vault models.

## Further Information

For detailed documentation, advanced configuration options, and troubleshooting, please refer to the full documentation available in the repository.

---

Developed with :black_heart: by Thomas F McGeehan V
