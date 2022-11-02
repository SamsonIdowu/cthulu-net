import json
import ipaddress
import uuid


class Task(dict):
    def __init__(self, ip, port, type, completed,
        completed_by, ip_status, port_status):
        super().__init__(self)
        if ip == None:
            self['task_id'] = None
        else:
            self['task_id'] = str(uuid.uuid4())
        self['ip'] = ip
        self['port'] = port
        self['type'] = type
        self['completed'] = completed
        self['completed_by'] = completed_by
        self['ip_status'] = ip_status
        self['port_status'] = port_status
    def toJSON(self):
        data = {
            'id': self['task_id'],
            'ip': self['ip'],
            'port': self['port'],
            'type': self['type']
        }
        return json.dumps(data)

class TaskPool(list):
    def __init__(self):
        self.curr = 0
        super().__init__(self)
    def enqueue(self, task):
        self.append(task)
    def dequeue(self):
        if self.curr < len(self):
            val = self[self.curr]
            self.curr += 1
            return val
        return Task(None, None, None, None, None, None, None)
    def complete(self, taskid, botid, ip_status, port_status):
        for task in self:
            if task["task_id"] == taskid:
                idx = self.index(task)
                self[idx]["completed"] = True
                self[idx]["completed_by"] = botid
                self[idx]["ip_status"] = ip_status
                self[idx]["port_status"] = port_status
                print("found the task")
                return json.dumps(self[idx])
        return json.dumps("{}")
        
class TaskRegistrar:
    def registerTasksForJob(self, job):
        global taskpool
        for ip in ipaddress.ip_network(job['netblock']).hosts():
            task = Task(str(ip), job['port'], job['type'], False, "", "", "")
            print(task)
            taskpool.append(task)

taskpool = TaskPool()
taskregistrar = TaskRegistrar()
