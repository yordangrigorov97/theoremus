from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    # ex: /polls/5/
    path('<str:fromTS>/<str:toTS>/hour', views.agghour, name='agghour'),
    path('<str:fromTS>/<str:toTS>/day', views.aggday, name='aggday'),
]
