import request from 'supertest';
import express from 'express';
import userRoutes from '../src/routes/userRoutes';
import '../src/config/env'; 


const app = express();
app.use(express.json());
app.use('/api/users', userRoutes);

describe('Auth Service Tests', () => {
  let testToken = '';
  const email = 'test@example.com';

  it('should register a new user', async () => {
    const res = await request(app)
      .post('/api/users/register')
      .send({
        name: 'Test User',
        email: email,
        password: 'password123'
      });
    expect(res.statusCode).toBe(201);
    expect(res.body.message).toBe('User registered successfully');
  });

  it('should not allow duplicate registration', async () => {
    const res = await request(app)
      .post('/api/users/register')
      .send({
        name: 'Test User',
        email: email,
        password: 'password123'
      });
    expect(res.statusCode).toBe(400);
    expect(res.body.message).toBe('User already exists');
  });

  it('should login with correct credentials', async () => {
    const res = await request(app)
      .post('/api/users/login')
      .send({
        email: email,
        password: 'password123'
      });
    expect(res.statusCode).toBe(200);
    expect(res.body.token).toBeDefined();
    testToken = res.body.token;
  });

  it('should fail login with wrong password', async () => {
    const res = await request(app)
      .post('/api/users/login')
      .send({
        email: email,
        password: 'wrongpass'
      });
    expect(res.statusCode).toBe(401);
  });

  it('should access protected /me route with token', async () => {
    const res = await request(app)
      .get('/api/users/me')
      .set('Authorization', `Bearer ${testToken}`);
    expect(res.statusCode).toBe(200);
    expect(res.body.user.email).toBe(email);
  });

  it('should reject /me without token', async () => {
    const res = await request(app)
      .get('/api/users/me');
    expect(res.statusCode).toBe(401);
  });
});
