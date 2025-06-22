import { Router } from 'express';
import { registerUser, loginUser } from '../controllers/userController';
import { authenticateToken } from '../middleware/authMiddleware';
import { getCurrentUser } from '../controllers/userController';
import { getAllUsers } from '../controllers/userController';
import { authorizeRole } from '../middleware/roleMiddleware';

const router = Router();

router.post('/register', registerUser);
router.post('/login', loginUser);

export default router;

router.get('/me', authenticateToken, getCurrentUser);

router.get('/all', authenticateToken, authorizeRole('admin'), getAllUsers);