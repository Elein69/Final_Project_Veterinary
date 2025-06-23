import dotenv from 'dotenv';
dotenv.config();

export const DB_CONFIG = {
  host: process.env.DB_HOST || 'host.docker.internal',
  port: Number(process.env.DB_PORT),
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
};

export const PORT = process.env.PORT || 3000;
export const JWT_SECRET = process.env.JWT_SECRET || 'default_secret';
export const IS_CI = process.env.CI === 'true';
