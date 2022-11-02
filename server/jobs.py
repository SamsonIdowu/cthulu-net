import uuid
import json

class Job(dict):
    def __init__(self, netblock, port, job_type):
        super().__init__(self)
        self['job_id'] = uuid.uuid4().hex
        self['netblock'] = netblock
        self['port'] = port
        self['type'] = job_type # scan/brute/spread/
    def toJSON(self):
        return json.dumps(self)