import express from 'express';
import cors from 'cors';
import dotenv from 'dotenv';
import userRoutes from './routes/userRoutes';
import { db } from './db';
import { createUserTable } from './models/userModel';

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());
app.use(cors());

app.use('/api/users', userRoutes);

app.get('/', (_req, res) => {
  res.send('Auth service is running!');
});

app.listen(PORT, () => {
  console.log(`Auth service listening on port ${PORT}`);
});

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