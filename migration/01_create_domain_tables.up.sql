CREATE TABLE organization
(
    inn          VARCHAR(12)  NOT NULL UNIQUE,
    name         VARCHAR(200) NOT NULL UNIQUE,
    exist_museum BOOLEAN DEFAULT FALSE
)

CREATE TABLE museum
(
    inn                                 VARCHAR(12)  NOT NULL REFERENCES organization (inn),
--     organization_name VARCHAR(200) NOT NULL,
    kpp                                 VARCHAR(100),
    founder                             VARCHAR(100),
    museum_activity_in_charter          BOOLEAN DEFAULT FALSE,
    name                                VARCHAR(200) NOT NULL,
    museum_legal_status                 VARCHAR(50)  NOT NULL,
    is_memorial_reserve_museum          BOOLEAN DEFAULT FALSE,
    is_historical_memorial_reserve      BOOLEAN DEFAULT FALSE,
    is_art_museum                       BOOLEAN DEFAULT FALSE,
    is_museum_reserve                   BOOLEAN DEFAULT FALSE,
    is_estate_museum                    BOOLEAN DEFAULT FALSE,
    is_palace_park_ensemble             BOOLEAN DEFAULT FALSE,
    is_historical_architectural_reserve BOOLEAN DEFAULT FALSE,
    annual_visitor_capacity             SMALLINT,
    internal_visitors_count             SMALLINT,
    external_visitors_count             SMALLINT,
    is_valuable_cultural_heritage       BOOLEAN DEFAULT FALSE,
    valuable_museum_items_count         SMALLINT
)

CREATE TABLE labor_resources
(
    inn                             VARCHAR(12) NOT NULL REFERENCES organization (inn),
    id                              SERIAL PRIMARY KEY,
    total_staff_annual              DECIMAL(10, 2),
    research_staff_internal         DECIMAL(4, 2),
    core_operational_staff_internal DECIMAL(4, 2),
    admin_support_staff_internal    DECIMAL(4, 2),
    research_staff_external         DECIMAL(4, 2),
    core_operational_staff_external DECIMAL(4, 2),
    admin_support_staff_external    DECIMAL(4, 2)
)

CREATE TABLE labor_resources_fot
(
    inn                             VARCHAR(12) NOT NULL REFERENCES organization (inn),
    labor_resources_id              INTEGER REFERENCES labor_resources (id) total_staff_annual              DECIMAL(12, 2),
    research_staff_internal         DECIMAL(10, 2),
    core_operational_staff_internal DECIMAL(10, 2),
    admin_support_staff_internal    DECIMAL(10, 2),
    research_staff_external         DECIMAL(10, 2),
    core_operational_staff_external DECIMAL(10, 2),
    admin_support_staff_external    DECIMAL(10, 2)
)

CREATE TABLE museum_activity_cost_distribution
(
    inn                                VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                               SMALLINT,
    access_via_stationary_exhibitions  INTEGER,
    access_via_mobile_exhibitions      INTEGER,
    collection_management_preservation INTEGER,
    restoration_conservation_works     INTEGER,
    public_display_stationary          INTEGER,
    public_display_mobile              INTEGER,
    public_display_online              INTEGER
)

CREATE TABLE another_activity_type
(
    inn  VARCHAR(12) NOT NULL REFERENCES organization (inn),
    name VARCHAR(200)
)

CREATE TABLE income_generating_activity
(
    inn                                VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                               SMALLINT,
    access_via_stationary_exhibitions  INTEGER,
    access_via_mobile_exhibitions      INTEGER,
    collection_management_preservation INTEGER,
    restoration_conservation_works     INTEGER,
    public_display_stationary          INTEGER,
    public_display_mobile              INTEGER,
    public_display_online              INTEGER
)

CREATE TABLE total_volume
(
    inn                        VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                       SMALLINT,
    exhibitions_total_count    INTEGER,
    museum_collections_count   INTEGER,
    individual_artifacts_count INTEGER,
    internal_audience_count    INTEGER,
    external_public_count      INTEGER
)

CREATE TABLE governmental_assignment_volume
(
    inn                        VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                       SMALLINT,
    exhibitions_total_count    INTEGER,
    museum_collections_count   INTEGER,
    individual_artifacts_count INTEGER,
    internal_audience_count    INTEGER,
    external_public_count      INTEGER
)

CREATE TABLE governmental_assignment_volume
(
    inn                        VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                       SMALLINT,
    exhibitions_total_count    INTEGER,
    museum_collections_count   INTEGER,
    individual_artifacts_count INTEGER,
    internal_audience_count    INTEGER,
    external_public_count      INTEGER
)

CREATE TABLE revenue_generating_activity_volume
(
    inn                        VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                       SMALLINT,
    exhibitions_total_count    INTEGER,
    museum_collections_count   INTEGER,
    individual_artifacts_count INTEGER,
    internal_audience_count    INTEGER,
    external_public_count      INTEGER
)

CREATE TABLE museum_revenues_expenses
(
    inn                          VARCHAR(12) NOT NULL REFERENCES organization (inn),
    year                         SMALLINT,
    total_revenue                DECIMAL(10, 2),
    state_assignment_subsidy     DECIMAL(8, 2),
    earned_revenue               DECIMAL(8, 2),
    other_funding_sources        DECIMAL(8, 2),

    total_expenses               DECIMAL(10, 2),
    inventory_assets_acquisition DECIMAL(8, 2),
    utility_services             DECIMAL(8, 2),
    communication_services       DECIMAL(8, 2),
    transportation_services      DECIMAL(8, 2),
    valuable_assets_acquisition  DECIMAL(8, 2),
    real_estate_maintenance      DECIMAL(8, 2),
    valuable_assets_maintenance  DECIMAL(8, 2),
    general_administrative_costs DECIMAL(8, 2),
    tax_payments                 DECIMAL(8, 2),
    other_expenses               DECIMAL(8, 2),
)


CREATE INDEX idx_organization_inn ON organization (inn);
CREATE INDEX idx_organization_exist_museum ON organization (exist_museum);

CREATE INDEX idx_museum_inn ON museum (inn);
CREATE INDEX idx_museum_name ON museum (name);
CREATE INDEX idx_museum_legal_status ON museum (museum_legal_status);
CREATE INDEX idx_museum_activity_charter ON museum (museum_activity_in_charter);
CREATE INDEX idx_museum_valuable_heritage ON museum (is_valuable_cultural_heritage);

CREATE INDEX idx_labor_resources_inn ON labor_resources (inn);

CREATE INDEX idx_labor_resources_fot_inn ON labor_resources_fot (inn);
CREATE INDEX idx_labor_resources_fot_labor_id ON labor_resources_fot (labor_resources_id);

CREATE INDEX idx_museum_activity_cost_inn ON museum_activity_cost_distribution (inn);
CREATE INDEX idx_income_generating_activity_inn ON income_generating_activity (inn);
CREATE INDEX idx_total_volume_inn ON total_volume (inn);
CREATE INDEX idx_governmental_assignment_volume_inn ON governmental_assignment_volume (inn);
CREATE INDEX idx_revenue_generating_activity_volume_inn ON revenue_generating_activity_volume (inn);
CREATE INDEX idx_museum_revenues_expenses_inn ON museum_revenues_expenses (inn);

CREATE INDEX idx_another_activity_type_inn ON another_activity_type (inn);
CREATE INDEX idx_another_activity_type_name ON another_activity_type (name);
count ON museum (valuable_museum_items_count);