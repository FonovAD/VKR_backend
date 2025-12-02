BEGIN;

/* Вузы */
CREATE TABLE IF NOT EXISTS organization (
    id           SERIAL PRIMARY KEY,
    inn          VARCHAR(12)  NOT NULL UNIQUE,
    name         VARCHAR(250) NOT NULL UNIQUE
);

/* Музеи */
CREATE TABLE IF NOT EXISTS museum (
    id                                  SERIAL PRIMARY KEY,
    organization_id                       INTEGER      NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
    inn                                 VARCHAR(12),
    kpp                                 TEXT,
    founder                             TEXT,
    museum_activity_in_charter          BOOLEAN      NOT NULL DEFAULT FALSE,
    name                                VARCHAR(300) NOT NULL,
    museum_legal_status                 TEXT         NOT NULL,
    is_memorial_reserve_museum          BOOLEAN      NOT NULL DEFAULT FALSE,
    is_historical_memorial_reserve      BOOLEAN      NOT NULL DEFAULT FALSE,
    is_art_museum                       BOOLEAN      NOT NULL DEFAULT FALSE,
    is_museum_reserve                   BOOLEAN      NOT NULL DEFAULT FALSE,
    is_estate_museum                    BOOLEAN      NOT NULL DEFAULT FALSE,
    is_palace_park_ensemble             BOOLEAN      NOT NULL DEFAULT FALSE,
    is_historical_architectural_reserve BOOLEAN      NOT NULL DEFAULT FALSE,
    annual_visitor_capacity             INTEGER,
    internal_visitors_count             INTEGER,
    external_visitors_count             INTEGER,
    is_valuable_cultural_heritage       BOOLEAN      NOT NULL DEFAULT FALSE,
    valuable_museum_items_count         INTEGER
);

CREATE INDEX IF NOT EXISTS idx_museum_organization ON museum (organization_id);

/* Трудовые ресурсы */
CREATE TABLE IF NOT EXISTS labor_resource (
    id                              SERIAL PRIMARY KEY,
    museum_id                       INTEGER     NOT NULL REFERENCES museum (id) ON DELETE CASCADE,
    total_staff_annual              DECIMAL(8, 2),
    research_staff_internal         DECIMAL(8, 2),
    core_operational_staff_internal DECIMAL(8, 2),
    admin_support_staff_internal    DECIMAL(8, 2),
    research_staff_external         DECIMAL(8, 2),
    core_operational_staff_external DECIMAL(8, 2),
    admin_support_staff_external    DECIMAL(8, 2)
);

CREATE TABLE IF NOT EXISTS labor_resource_fot (
    labor_resources_id              INTEGER REFERENCES labor_resource (id) ON DELETE CASCADE,
    total_staff_annual              DECIMAL(12, 2),
    research_staff_internal         DECIMAL(12, 2),
    core_operational_staff_internal DECIMAL(12, 2),
    admin_support_staff_internal    DECIMAL(12, 2),
    research_staff_external         DECIMAL(12, 2),
    core_operational_staff_external DECIMAL(12, 2),
    admin_support_staff_external    DECIMAL(12, 2)
);

CREATE INDEX IF NOT EXISTS idx_labor_resource_museum ON labor_resource (museum_id);

/* Типы деятельности */
CREATE TABLE IF NOT EXISTS activity_types (
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);


/* Справочник показателей объема */
CREATE TABLE IF NOT EXISTS volume_indicators (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

/* Основная таблица деятельности */
CREATE TABLE IF NOT EXISTS museum_activities (
    id                      BIGSERIAL PRIMARY KEY,
    museum_id               INTEGER  NOT NULL REFERENCES museum (id) ON DELETE CASCADE,
    activity_type_id        INTEGER REFERENCES activity_types (id),
    volume_indicator_id     INTEGER REFERENCES volume_indicators (id),
    visitor_category        TEXT     NOT NULL CHECK (visitor_category IN ('internal', 'external')),
    cost_share_percent      NUMERIC(5, 2),
    revenue_amount          NUMERIC(18, 2),
    total_count             NUMERIC(15, 0),
    state_task_count        NUMERIC(15, 0),
    revenue_activity_count  NUMERIC(15, 0),
    year                    SMALLINT NOT NULL DEFAULT 2022
);

/* Обеспечиваем уникальность комбинаций */
CREATE UNIQUE INDEX IF NOT EXISTS museum_activities_unique_standard
    ON museum_activities (museum_id, activity_type_id, visitor_category, year)
    WHERE activity_type_id IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_museum_activities_museum ON museum_activities (museum_id);
CREATE INDEX IF NOT EXISTS idx_museum_activities_volume_indicator ON museum_activities (volume_indicator_id);

/* Доходы и расходы */
CREATE TABLE IF NOT EXISTS museum_revenues_expense (
    id                           SERIAL PRIMARY KEY,
    museum_id                    INTEGER NOT NULL REFERENCES museum (id) ON DELETE CASCADE,
    year                         SMALLINT NOT NULL,
    total_revenue                DECIMAL(12, 2),
    state_assignment_subsidy     DECIMAL(12, 2),
    earned_revenue               DECIMAL(12, 2),
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
    UNIQUE (museum_id, year)
);

/* Другие источники финансирования */
CREATE TABLE IF NOT EXISTS other_funding_sources (
    id                         SERIAL PRIMARY KEY,
    museum_revenues_expense_id INTEGER NOT NULL REFERENCES museum_revenues_expense (id) ON DELETE CASCADE,
    source_name                TEXT,
    amount                     DECIMAL(12, 2) NOT NULL
);

/* Другие расходы */
CREATE TABLE IF NOT EXISTS other_expenses (
    id                         SERIAL PRIMARY KEY,
    museum_revenues_expense_id INTEGER NOT NULL REFERENCES museum_revenues_expense (id) ON DELETE CASCADE,
    expenses_name              TEXT,
    amount                     DECIMAL(12, 2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_other_funding_sources_revenue_expense_id ON other_funding_sources (museum_revenues_expense_id);
CREATE INDEX IF NOT EXISTS idx_other_expenses_revenue_expense_id ON other_expenses (museum_revenues_expense_id);

/* Заполняем справочник показателей объема */
INSERT INTO volume_indicators (id, name)
VALUES
    (1, 'Количество экспозиций (выставок)'),
    (2, 'Количество музейных предметов и музейных коллекций'),
    (3, 'Количество предметов'),
    (4, 'Количество посетителей')
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name;

/* Заполняем справочник типовых видов деятельности */
INSERT INTO activity_types (id, name)
VALUES
    (1, 'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (в стационарных условиях)'),
    (2, 'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (вне стационара)'),
    (3, 'Комплектование, учет, обеспечение безопасности и сохранности музейных предметов и музейных коллекций'),
    (4, 'Проведение реставрационных работ в отношении музейных предметов и музейных коллекций'),
    (5, 'Публичный показ музейных предметов, музейных коллекций (в стационарных условиях)'),
    (6, 'Публичный показ музейных предметов, музейных коллекций (вне стационара)'),
    (7, 'Публичный показ музейных предметов, музейных коллекций (удаленно через интернет)')
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name;

COMMIT;
