from fastapi import FastAPI
from time import sleep
from fastapi.middleware.cors import CORSMiddleware


busy_duration: int = 5  
busy: bool = False

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/message")
def read_root():
    return {"message": "Hello World from server!"}

@app.get("/health")
def health_check():
    return "ok"

@app.post("/task")
async def process_task(task: dict = {}):
    sleep(busy_duration) 
    return {"status": "processed", "task": task}