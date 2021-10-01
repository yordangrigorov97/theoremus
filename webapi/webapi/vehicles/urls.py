from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    # ex: /polls/5/
    path('<int:fromTS>/<int:toTS>/hour', views.agghour, name='agghour'),
    path('<int:fromTS>/<int:toTS>/day', views.aggday, name='aggday'),
]
