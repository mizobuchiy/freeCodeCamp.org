from rest_framework import serializer
from .models import Product


class ProductSerializer(serializer):
    class Meta:
        model = Product
        fields = '__all__'
