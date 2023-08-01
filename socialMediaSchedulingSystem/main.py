import time
from abc import ABCMeta, abstractclassmethod


class Post:
    def __init__(self, post_id, content, schedule):
        self.post_id = post_id
        self.content = content
        self.schedule = schedule


class Command(ABCMeta):

    @abstractclassmethod
    def execute(self):
        pass


class PostScheduler:
    def __init__(self):
        self.posts = []
        self.observers = []
        self.cancel_flag = False

    def add_observer(self, observer):
        self.observers.append(observer)

    def schedule_post(self, post):
        self.posts.append(post)

    def cancel_post(self, post_id):
        self.posts = [post for post in self.posts if post.post_id != post_id]

    def start(self):
        while not self.cancel_flag:
            for post in self.posts:
                if time.time() >= post.schedule:
                    for observer in self.observers:
                        observer.update(post)
                    self.posts.remove(post)
            time.sleep(1)


class SchedulePostCommand(Command):
    def __init__(self, post_scheduler, post):
        self.post_scheduler = post_scheduler
        self.post = post

    def execute(self):
        self.post_scheduler.schedule_post(self.post)
        print(f"投稿がスケジュールされました。ID: {self.post.post_id}, 内容: {self.post.content}, スケジュール日時: {self.post.schedule}")


class CancelPostCommand(Command):
    def __init__(self, post_scheduler, post_id):
        self.post_scheduler = post_scheduler
        self.post_id = post_id

    def execute(self):
        self.post_scheduler.cancel_post(self.post_id)
        print(f"投稿がキャンセルされました。ID: {self.post_id}")


class Observer(ABCMeta):

    @abstractclassmethod
    def update(self, post):
        pass


class PostObserver(Observer):
    def update(self, post):
        print(f"投稿が実行されました。ID: {post.post_id}, 内容: {post.content}")


if __name__ == "__main__":
    post_scheduler = PostScheduler()

    post_observer = PostObserver()
    post_scheduler.add_observer(post_observer)

    post1 = Post(1, "Hello, World!", time.time() + 10)
    post2 = Post(2, "This is a scheduled post.", time.time() + 20)
    SchedulePostCommand(post_scheduler, post1).execute()
    SchedulePostCommand(post_scheduler, post2).execute()

    time.sleep(5)
    CancelPostCommand(post_scheduler, 1).execute()

    time.sleep(25)
    post_scheduler.cancel_flag = True
