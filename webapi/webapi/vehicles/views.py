from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
from .vehicles_agg import aggregate


def index(request):
    return HttpResponse("Hello, world. Let's count vehicles!")

def agghour(request, fromTS, toTS):
    result = aggregate(fromTS, toTS, 'hour')
    # for item in item_details:
    #     # This does not give a very readable output
    #     print(item)
    return HttpResponse(result)

def aggday(request, fromTS, toTS):
    result = aggregate(fromTS, toTS, 'day')
    # for item in item_details:
    #     # This does not give a very readable output
    #     print(item)
    return HttpResponse(result)
