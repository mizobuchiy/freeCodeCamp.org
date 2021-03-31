import pika

params = pika.URLParameters(
    'amqps://puaiumzk:UZPxGrIIoIgQOu2UrQ3k6vkGU7YMb2U6@hornet.rmq.cloudamqp.com/puaiumzk')

connection = pika.BlockingConnection(params)

channel = connection.channel()


def publish():
    channel.basic_publish(exchange='', routing_key='main', body='hello')
