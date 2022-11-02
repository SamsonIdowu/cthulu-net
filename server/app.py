from crypt import methods
from flask import Flask, request
import json

from bot import botpool
from task import taskpool, taskregistrar
from jobs import Job

app = Flask(__name__)

# BotRegistrar
@app.route("/bot/", methods=["POST"])
def register():
    data = request.get_json()
    data['ip'] = request.remote_addr
    data['ram'] = data['ram'] / 1024
    botid = botpool.addBot(data)
    print(botpool)
    return f"{botid}"

# BotTasker
@app.route("/bot/tasks/", methods=["GET"])
def give_task():
    return taskpool.dequeue().toJSON()


# BotReportsCollector
@app.route("/bot/<botid>", methods=["PUT"])
def report_task(botid):
    data = request.get_json()
    taskid = data["TaskId"]
    ip_status = data["Ip_status"]
    port_status = data["Port_status"]
    return taskpool.complete(taskid, botid, ip_status, port_status)

# JobRegistrar
@app.route("/operator/jobs", methods=["POST"])
def register_job():
    data = request.get_json()
    job = Job(data['netblock'], data['port'], data['type'])
    taskregistrar.registerTasksForJob(job)
    return job.toJSON()

# BotPoolStats
@app.route("/operator/bots", methods=["GET"])
def getbot_pool():
    bots = botpool.getAllBots()
    payload = []
    for botid in bots:
        bot = botpool.getBot(botid)
        bot["id"] = botid
        payload.append(bot)
    return json.dumps(payload)

# JobCompletionStat
@app.route("/operator/jobs", methods=["GET"])
def getjob_stats():
    tasks_completed = []
    for task in taskpool:
        if task["completed"] == True:
            tasks_completed.append(task)
    payload = {
        "total_tasks": len(taskpool),
        "total_completed_tasks": len(tasks_completed),
        "tasks_completed": tasks_completed
    }
    return json.dumps(payload)


# TaskCompletionPerBot
#@app.route("/operator/tasks", methods=["GET"])
#def getbot_tasks():
#    payload = {}
#    for botid in botpool.getAllBots():
#        payload[botid] = []
#        for task in taskpool:
#            if (task["completed"] == True) and (task['completed_by'] == botid):
#                payload[botid].append(task)
#    return json.dumps(payload)

@app.route("/operator/tasks", methods=["GET"])
def getall_tasks():
    payload = []
    for task in taskpool:
        payload.append(task)
    return json.dumps(payload)

@app.route("/operator/tasks/summary", methods=["GET"])
def tasks_summary():
    payload = { "total_tasks": 0, "tasks_completed": 0, "tasks_remain": 0}
    n = 0
    for task in taskpool:
        if task["completed"] == True:
            n += 1
    payload["total_tasks"] = len(taskpool)
    payload["tasks_completed"] = n
    payload["tasks_remain"] = len(taskpool) - n
    return json.dumps(payload)


@app.route("/operator/tasks/hosts/summary", methods=["GET"])
def hosts_status():
    payload = {"up": 0, "down": 0}
    up = 0
    n = 0
    for task in taskpool:
        if task["completed"] == True:
            if task["ip_status"] == "up":
                up += 1
            n += 1
    payload["up"] = up
    payload["down"] = n - up
    return json.dumps(payload)
