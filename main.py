import asyncio
import logging
from os import path
from subprocess import Popen

logging.basicConfig(filename="/tmp/decktweaks.log",
                    format='[DeckTweaks] %(asctime)s %(levelname)s %(message)s',
                    filemode='w+',
                    force=True)
logger=logging.getLogger()
logger.setLevel(logging.INFO)

PROJECT_DIR = path.realpath(path.dirname(__file__))
BACKEND_STDOUT_FILE = "/tmp/decktweaks_backend-stdout.log"
BACKEND_STDERR_FILE = "/tmp/decktweaks_backend-stderr.log"

class Plugin:
    async def _main(self):
        logger.info("Starting Plugin")

        # Start the backend server, directing outputs to a log file.
        backend_server_bin = path.join(PROJECT_DIR, "bin/server")
        backend_stdout_fd = open(BACKEND_STDOUT_FILE, 'a')
        backend_stderr_fd = open(BACKEND_STDERR_FILE, 'a')

        logger.info(f"Starting backend {backend_server_bin}")
        Popen(
            [backend_server_bin, "start"],
            stdout=backend_stdout_fd,
            stderr=backend_stderr_fd,
            cwd=PROJECT_DIR,
        )

        while True:
            await asyncio.sleep(1)
