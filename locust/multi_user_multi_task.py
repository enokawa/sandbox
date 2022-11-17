from locust import HttpUser, SequentialTaskSet, User, constant, task

class ForumSection(SequentialTaskSet):
    @task()
    def view_thread(self):
        print("view_thread")
        self.client.get("/")
        pass

    @task()
    def create_thread(self):
        print("create_thread")
        pass

class AnotherSection(SequentialTaskSet):
    @task()
    def delete_thread(self):
        print("delete_thread")
        self.client.get("/")
        pass

    @task()
    def update_thread(self):
        print("update_thread")
        pass

class LoggedInUser(HttpUser):
    host = "https://enokawa.dev"
    wait_time = constant(1)
    tasks = [ForumSection]

class AnotherUser(HttpUser):
    host = "https://enokawa.dev"
    wait_time = constant(1)
    tasks = [AnotherSection]

class StressUser(User):
  host = "https://enokawa.dev"
  tasks = {LoggedInUser: 1, AnotherUser: 1}
