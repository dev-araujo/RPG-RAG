import express from 'express';
import cors from 'cors';
import dotenv from 'dotenv';
import askRouter from './features/ask/ask.routes';

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;

app.use(cors({ origin: 'http://localhost:4200' }));
app.use(express.json());

app.get('/health', (_req, res) => {
  res.json({
    status: 'ok',
    service: 'LoreKeeper Backend',
    timestamp: new Date().toISOString(),
  });
});

app.use('/api/ask', askRouter);

app.listen(PORT, () => {
  console.log(`\nLoreKeeper Backend running on http://localhost:${PORT}`);
  console.log(`   RAG Engine: ${process.env.RAG_ENGINE_ADDRESS || 'localhost:50051'}`);
  console.log(`   Ready to serve the Dungeon Master!\n`);
});

export default app;
