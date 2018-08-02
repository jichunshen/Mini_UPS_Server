from django.conf.urls import url
from . import views

app_name = 'event'

urlpatterns = [
    url(r'^$', views.index, name='index'),
    url(r'^register/$', views.register, name='register'),
    url(r'^login_user/$', views.login_user, name='login_user'),
    url(r'^track_id/$', views.track_id, name='track_id'),
    url(r'^logout_user/$', views.logout_user, name='logout_user'),
    url(r'^(?P<order_id>[0-9]+)/$', views.detail, name='detail'),
    url(r'^(?P<order_id>[0-9]+)/$', views.change_destination, name='change_destination'),
    # url(r'^(?P<event_id>[0-9]+)/$', views.detail, name='detail'),
]