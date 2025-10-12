-- Удаление индексов
DROP INDEX IF EXISTS idx_another_activity_type_name;
DROP INDEX IF EXISTS idx_another_activity_type_inn;
DROP INDEX IF EXISTS idx_museum_revenues_expenses_inn;
DROP INDEX IF EXISTS idx_revenue_generating_activity_volume_inn;
DROP INDEX IF EXISTS idx_governmental_assignment_volume_inn;
DROP INDEX IF EXISTS idx_total_volume_inn;
DROP INDEX IF EXISTS idx_income_generating_activity_inn;
DROP INDEX IF EXISTS idx_museum_activity_cost_inn;
DROP INDEX IF EXISTS idx_labor_resources_fot_labor_id;
DROP INDEX IF EXISTS idx_labor_resources_fot_inn;
DROP INDEX IF EXISTS idx_labor_resources_inn;
DROP INDEX IF EXISTS idx_museum_valuable_heritage;
DROP INDEX IF EXISTS idx_museum_activity_charter;
DROP INDEX IF EXISTS idx_museum_legal_status;
DROP INDEX IF EXISTS idx_museum_name;
DROP INDEX IF EXISTS idx_museum_inn;
DROP INDEX IF EXISTS idx_organization_exist_museum;
DROP INDEX IF EXISTS idx_organization_inn;

-- Удаление таблиц в правильном порядке (сначала дочерние, затем родительские)
DROP TABLE IF EXISTS museum_revenues_expenses;
DROP TABLE IF EXISTS revenue_generating_activity_volume;
DROP TABLE IF EXISTS governmental_assignment_volume;
DROP TABLE IF EXISTS total_volume;
DROP TABLE IF EXISTS income_generating_activity;
DROP TABLE IF EXISTS museum_activity_cost_distribution;
DROP TABLE IF EXISTS another_activity_type;
DROP TABLE IF EXISTS labor_resources_fot;
DROP TABLE IF EXISTS labor_resources;
DROP TABLE IF EXISTS museum;
DROP TABLE IF EXISTS organization;