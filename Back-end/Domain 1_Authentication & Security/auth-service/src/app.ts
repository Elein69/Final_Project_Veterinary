import { db } from './db';
import { IS_CI } from './config/env';
import { createUserTable } from './models/userModel';

if (!IS_CI) {
  db.connect()
    .then(() => {
      console.log('✅ Connected to PostgreSQL database');
      return createUserTable();
    })
    .then(() => {
      console.log('✅ Users table ready');
    })
    .catch((err: any) => {
      console.error('❌ Database connection error:', err);
    });
} else {
  console.log('⚠️ Skipping DB connection and table creation in CI');
}
