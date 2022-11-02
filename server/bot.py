import hashlib
import time


class Bot(dict):
    def __init__(self, hostname, username, ncpu, ram, ip, last_seen=None):
        super().__init__(self)
        self["hostname"] = hostname
        self["username"] = username
        self["ncpu"] = ncpu
        self["ram"] = ram
        self["ip"] = ip
        self["last_seen"] = last_seen

class BotPool(dict):
    def getBot(self, uuid):
        return self[uuid]
    def getAllBots(self):
        return self.keys()
    def addBot(self, data):
        combo = data["hostname"] + data["username"]
        botid = hashlib.md5(combo.encode('utf-8')).hexdigest()
        if botid not in self:
            self[botid] = Bot(
                data["hostname"],
                data["username"],
                data["ncpu"],
                data["ram"],
                data["ip"]
            )
        self[botid]["last_seen"] = time.time()
        return botid

botpool = BotPool()
