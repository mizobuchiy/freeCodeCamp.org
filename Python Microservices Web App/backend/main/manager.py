from main import app, db
from flask_migrate import Migrate, MigrateCommand
from flask_script import Manager

migrate = Migrate(app, db)

maneger = Manager(app)
maneger.add_command('db', MigrateCommand)

if __name__ == '__main__':
    maneger.run()
