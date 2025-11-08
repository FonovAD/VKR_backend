CREATE TABLE organization
(
    id           SERIAL PRIMARY KEY,
    inn          VARCHAR(12)  NOT NULL UNIQUE,
    name         VARCHAR(250) NOT NULL UNIQUE,
    exist_museum BOOLEAN DEFAULT FALSE
);

CREATE TABLE museum
(
    id_owner                            INTEGER      NOT NULL REFERENCES organization (id),
    id                                  SERIAL,
    inn                                 VARCHAR(12),
    kpp                                 TEXT,
    founder                             TEXT,
    museum_activity_in_charter          BOOLEAN DEFAULT FALSE,
    name                                VARCHAR(300) NOT NULL,
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
    labor_resources_id              INTEGER REFERENCES labor_resource (id),
    total_staff_annual              DECIMAL(12, 2),
    research_staff_internal         DECIMAL(12, 2),
    core_operational_staff_internal DECIMAL(12, 2),
    admin_support_staff_internal    DECIMAL(12, 2),
    research_staff_external         DECIMAL(12, 2),
    core_operational_staff_external DECIMAL(12, 2),
    admin_support_staff_external    DECIMAL(12, 2)
);

CREATE TABLE activity_types
(
    id               INTEGER PRIMARY KEY,
    name             TEXT NOT NULL,
    volume_indicator TEXT NOT NULL -- Описание показателя объема
);

INSERT INTO activity_types (id, name, volume_indicator)
VALUES (1,
        'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (в стационарных условиях)',
        'Количество экспозиций (выставок)'),
       (2,
        'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (вне стационара)',
        'Количество экспозиций (выставок)'),
       (3, 'Комплектование, учет, обеспечение безопасности и сохранности музейных предметов и музейных коллекций',
        'Количество музейных предметов и музейных коллекций'),
       (4, 'Проведение реставрационных работ в отношении музейных предметов и музейных коллекций',
        'Количество предметов'),
       (5, 'Публичный показ музейных предметов, музейных коллекций (в стационарных условиях)',
        'Количество посетителей'),
       (6, 'Публичный показ музейных предметов, музейных коллекций (вне стационара)', 'Количество посетителей'),
       (7, 'Публичный показ музейных предметов, музейных коллекций (удаленно через интернет)',
        'Количество посетителей');
)

CREATE TABLE museum_activities
(
    inn                    VARCHAR(12) NOT NULL,
    activity_type_id       INTEGER     NOT NULL REFERENCES activity_types (id),
    visitor_category       TEXT        NOT NULL CHECK (visitor_category IN ('internal', 'external')), -- внутренние / сторонние
    cost_share_percent     NUMERIC(5, 2),
    revenue_amount         NUMERIC(18, 2),
    total_count            NUMERIC(15, 0),
    state_task_count       NUMERIC(15, 0),
    revenue_activity_count NUMERIC(15, 0),
    year                   SMALLINT    NOT NULL DEFAULT 2022,

    PRIMARY KEY (inn, activity_type_id, visitor_category, year),
);

CREATE INDEX idx_museum_activities_inn ON museum_activities (inn);

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
