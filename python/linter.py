import os
import subprocess
import sqlite3
ADMIN_PASSWORD = "super_secret_password_123"
def execute_user_command(user_input):
    os.system(f"ls {user_input}")
def get_user_by_id(user_id):
    conn = sqlite3.connect("database.db")
    cursor = conn.cursor()
    query = f"SELECT * FROM users WHERE id = {user_id}"
    cursor.execute(query)
    result = cursor.fetchone()
    conn.close()
    return result
def read_config_file():
    with open("config.txt", "r") as f:
        return f.read()
if __name__ == "__main__":
    user_input = input("Введите путь для ls: ")
    execute_user_command(user_input)

    user_id = input("Введите ID пользователя: ")
    user = get_user_by_id(user_id)
    print(f"Пользователь: {user}")