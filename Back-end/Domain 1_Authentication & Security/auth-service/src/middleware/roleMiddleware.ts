import { Request, Response, NextFunction, RequestHandler } from 'express';

export const authorizeRole = (roleRequired: string): RequestHandler => {
  return (req: Request, res: Response, next: NextFunction): void => {
    const user = (req as any).user;

    if (!user || user.role !== roleRequired) {
      res.status(403).json({ message: 'Forbidden: insufficient role' });
      return;
    }

    next();
  };
};

