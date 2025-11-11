BEGIN;

/* Организации */
CREATE TABLE IF NOT EXISTS organization (
    id           SERIAL PRIMARY KEY,
    inn          VARCHAR(12)  NOT NULL UNIQUE,
    name         VARCHAR(250) NOT NULL UNIQUE,
    exist_museum BOOLEAN      NOT NULL DEFAULT FALSE
);

/* Музеи */
CREATE TABLE IF NOT EXISTS museum (
    id                                  SERIAL PRIMARY KEY,
    id_owner                            INTEGER      NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
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

CREATE INDEX IF NOT EXISTS idx_museum_owner ON museum (id_owner);

/* Трудовые ресурсы */
CREATE TABLE IF NOT EXISTS labor_resource (
    id                              SERIAL PRIMARY KEY,
    id_owner                        INTEGER     NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
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

CREATE INDEX IF NOT EXISTS idx_labor_resource_owner ON labor_resource (id_owner);

/* Типы деятельности */
CREATE TABLE IF NOT EXISTS activity_types (
    id               INTEGER PRIMARY KEY,
    name             TEXT NOT NULL UNIQUE,
    volume_indicator TEXT NOT NULL
);

/* Пользовательские виды деятельности */
CREATE TABLE IF NOT EXISTS activity_custom_types (
    id               SERIAL PRIMARY KEY,
    name             TEXT NOT NULL UNIQUE,
    volume_indicator TEXT
);

/* Основная таблица деятельности */
CREATE TABLE IF NOT EXISTS museum_activities (
    id                      BIGSERIAL PRIMARY KEY,
    id_owner                INTEGER  NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
    activity_type_id        INTEGER REFERENCES activity_types (id),
    custom_activity_id      INTEGER REFERENCES activity_custom_types (id),
    visitor_category        TEXT     NOT NULL CHECK (visitor_category IN ('internal', 'external')),
    cost_share_percent      NUMERIC(5, 2),
    revenue_amount          NUMERIC(18, 2),
    total_count             NUMERIC(15, 0),
    state_task_count        NUMERIC(15, 0),
    revenue_activity_count  NUMERIC(15, 0),
    year                    SMALLINT NOT NULL DEFAULT 2022,
    CONSTRAINT museum_activities_activity_choice_chk CHECK (
        (activity_type_id IS NOT NULL AND custom_activity_id IS NULL)
        OR (activity_type_id IS NULL AND custom_activity_id IS NOT NULL)
    )
);

/* Обеспечиваем уникальность комбинаций */
CREATE UNIQUE INDEX IF NOT EXISTS museum_activities_unique_standard
    ON museum_activities (id_owner, activity_type_id, visitor_category, year)
    WHERE activity_type_id IS NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS museum_activities_unique_custom
    ON museum_activities (id_owner, custom_activity_id, visitor_category, year)
    WHERE custom_activity_id IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_museum_activities_owner ON museum_activities (id_owner);
CREATE INDEX IF NOT EXISTS idx_museum_activities_custom ON museum_activities (custom_activity_id);

/* Доходы и расходы */
CREATE TABLE IF NOT EXISTS museum_revenues_expense (
    id_owner                     INTEGER NOT NULL REFERENCES organization (id) ON DELETE CASCADE,
    year                         SMALLINT NOT NULL,
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
    other_expenses               DECIMAL(12, 2),
    PRIMARY KEY (id_owner, year)
);

/* Заполняем справочник типовых видов деятельности */
INSERT INTO activity_types (id, name, volume_indicator)
VALUES
    (1, 'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (в стационарных условиях)', 'Количество экспозиций (выставок)'),
    (2, 'Обеспечение доступа граждан к музейным предметам и музейным коллекциям путем создания экспозиций (выставок) (вне стационара)', 'Количество экспозиций (выставок)'),
    (3, 'Комплектование, учет, обеспечение безопасности и сохранности музейных предметов и музейных коллекций', 'Количество музейных предметов и музейных коллекций'),
    (4, 'Проведение реставрационных работ в отношении музейных предметов и музейных коллекций', 'Количество предметов'),
    (5, 'Публичный показ музейных предметов, музейных коллекций (в стационарных условиях)', 'Количество посетителей'),
    (6, 'Публичный показ музейных предметов, музейных коллекций (вне стационара)', 'Количество посетителей'),
    (7, 'Публичный показ музейных предметов, музейных коллекций (удаленно через интернет)', 'Количество посетителей')
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name,
    volume_indicator = EXCLUDED.volume_indicator;

COMMIT;
