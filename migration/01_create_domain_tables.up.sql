CREATE TABLE organization
(
    id           SERIAL      NOT NULL UNIQUE,
    inn          VARCHAR(12)  NOT NULL UNIQUE,
    name         VARCHAR(250) NOT NULL UNIQUE,
    exist_museum BOOLEAN DEFAULT FALSE
);

CREATE TABLE museum
(
    id_owner                            INTEGER      NOT NULL REFERENCES organization (id),
    inn                                 VARCHAR(12)  NOT NULL,
    kpp                                 TEXT,
    founder                             TEXT,
    museum_activity_in_charter          BOOLEAN DEFAULT FALSE,
    name                                VARCHAR(250) NOT NULL,
    museum_legal_status                 TEXT,
    is_memorial_reserve_museum          BOOLEAN DEFAULT FALSE,
    is_historical_memorial_reserve      BOOLEAN DEFAULT FALSE,
    is_art_museum                       BOOLEAN DEFAULT FALSE,
    is_museum_reserve                   BOOLEAN DEFAULT FALSE,
    is_estate_museum                    BOOLEAN DEFAULT FALSE,
    is_palace_park_ensemble             BOOLEAN DEFAULT FALSE,
    is_historical_architectural_reserve BOOLEAN DEFAULT FALSE,
    annual_visitor_capacity             INTEGER,
    internal_visitors_count             INTEGER,
    external_visitors_count             INTEGER,
    is_valuable_cultural_heritage       BOOLEAN DEFAULT FALSE,
    valuable_museum_items_count         INTEGER
);

CREATE TABLE labor_resource
(
    id_owner                        INTEGER NOT NULL REFERENCES organization (id),
    id                              SERIAL PRIMARY KEY,
    total_staff_annual              DECIMAL(8, 2),
    research_staff_internal         DECIMAL(8, 2),
    core_operational_staff_internal DECIMAL(8, 2),
    admin_support_staff_internal    DECIMAL(8, 2),
    research_staff_external         DECIMAL(8, 2),
    core_operational_staff_external DECIMAL(8, 2),
    admin_support_staff_external    DECIMAL(8, 2)
);

CREATE TABLE labor_resource_fot
(
    labor_resources_id              INTEGER REFERENCES labor_resources (id),
    total_staff_annual              DECIMAL(12, 2),
    research_staff_internal         DECIMAL(12, 2),
    core_operational_staff_internal DECIMAL(12, 2),
    admin_support_staff_internal    DECIMAL(12, 2),
    research_staff_external         DECIMAL(12, 2),
    core_operational_staff_external DECIMAL(12, 2),
    admin_support_staff_external    DECIMAL(12, 2)
);

CREATE TABLE activity
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(500) NOT NULL UNIQUE,
    location_type VARCHAR(50) CHECK (location_type IN ('internal', 'external', 'online'))
);

CREATE TABLE metric
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(300) NOT NULL UNIQUE,
    audience_type VARCHAR(50)  NOT NULL DEFAULT 'all'
        CHECK (audience_type IN ('internal', 'external', 'all'))
);

CREATE TABLE activity_record
(
    id          SERIAL PRIMARY KEY,
    activity_id INTEGER NOT NULL REFERENCES activities (id) ON DELETE CASCADE,
    metric_id   INTEGER NOT NULL REFERENCES metrics (id) ON DELETE CASCADE,
    cost_share  DECIMAL(5, 2)
);

CREATE TABLE funding_source
(
    id                      SERIAL PRIMARY KEY,
    record_id               INTEGER NOT NULL REFERENCES activity_records (id) ON DELETE CASCADE,
    revenue                 DECIMAL(15, 2),
    total_volume            INTEGER,
    state_assignment_volume INTEGER,
    income_activity_volume  INTEGER
);

CREATE TABLE museum_revenues_expense
(
    id_owner                     INTEGER NOT NULL REFERENCES organization (id),
    year                         SMALLINT,
    total_revenue                DECIMAL(12, 2),
    state_assignment_subsidy     DECIMAL(12, 2),
    earned_revenue               DECIMAL(12, 2),
    other_funding_sources        DECIMAL(12, 2),

    total_expenses               DECIMAL(12, 2),
    inventory_assets_acquisition DECIMAL(12, 2),
    utility_services             DECIMAL(12, 2),
    communication_services       DECIMAL(12, 2),
    transportation_services      DECIMAL(12, 2),
    valuable_assets_acquisition  DECIMAL(12, 2),
    real_estate_maintenance      DECIMAL(12, 2),
    valuable_assets_maintenance  DECIMAL(12, 2),
    general_administrative_costs DECIMAL(12, 2),
    tax_payments                 DECIMAL(12, 2),
    other_expenses               DECIMAL(12, 2)
);
