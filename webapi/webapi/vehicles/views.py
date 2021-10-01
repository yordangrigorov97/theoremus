from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse


def index(request):
    return HttpResponse("Hello, world. Let's count vehicles!")

def agghour(request, fromTS, toTS):
    return HttpResponse("hour")

def aggday(request, fromTS, toTS):
    return HttpResponse("day")
