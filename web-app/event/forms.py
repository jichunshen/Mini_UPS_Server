from django import forms
from django.contrib.auth.models import User
from .models import Orders

# from .models import Event


# class EventForm(forms.ModelForm):
#
#     class Meta:
#         model = Event
#         fields = ['event_name', 'event_time', 'event_owner', 'location']


class UserForm(forms.ModelForm):
    password = forms.CharField(widget=forms.PasswordInput)

    class Meta:
        model = User
        fields = ['username', 'email', 'password']

class DestinationForm(forms.ModelForm):
    des_x = forms.IntegerField
    des_y = forms.IntegerField

    class Meta:
        model = Orders
        fields = ['des_x', 'des_y']