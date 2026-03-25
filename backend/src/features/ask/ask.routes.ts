import { Router } from "express";
import { askController } from "./ask.controller";

const router = Router();

router.post("/", askController);

export default router;
