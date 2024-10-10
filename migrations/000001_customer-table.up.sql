-- Staff schema
CREATE SCHEMA IF NOT EXISTS staff;
-- Customer
CREATE TABLE IF NOT EXISTS staff.customer (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(250) NOT NULL,
  created TIMESTAMP(0) WITH TIME ZONE NOT NULL  DEFAULT Now(),
  updated TIMESTAMP(0) WITH TIME ZONE NOT NULL  DEFAULT Now(),
  active BOOLEAN NOT NULL DEFAULT true,
  version INTEGER NOT NULL DEFAULT 1
);

COMMENT ON TABLE staff.customer IS 'Customer profile';
COMMENT ON COLUMN staff.customer.version IS 'OCC';