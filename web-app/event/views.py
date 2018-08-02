from django.contrib.auth import authenticate, login
from django.contrib.auth import logout
from django.http import JsonResponse
from django.shortcuts import render, get_object_or_404, redirect
from django.db.models import Q
from django.urls import reverse
from .forms import UserForm, DestinationForm
from .models import Orders, Item



def index(request):
    if not request.user.is_authenticated:
        return render(request, 'event/login.html')
    else:
        orders = Orders.objects.filter(user_name = request.user.username)
        # return render(request, 'event/index.html', {'events': events})
        return render(request, 'event/index.html', {'orders': orders})

def logout_user(request):
    logout(request)
    form = UserForm(request.POST or None)
    context = {
        "form": form,
    }
    return render(request, 'event/login.html', context)

def login_user(request):
    if request.method == "POST":
        username = request.POST['username']
        password = request.POST['password']
        user = authenticate(username=username, password=password)
        if user is not None:
            if user.is_active:
                login(request, user)
                orders = Orders.objects.filter(user_name = request.user.username)
                return render(request, 'event/index.html', {'orders': orders})
            else:
                return render(request, 'event/login.html', {'error_message': 'Your account has been disabled'})
        else:
            return render(request, 'event/login.html', {'error_message': 'Invalid login'})
    return render(request, 'event/login.html')

def register(request):
    form = UserForm(request.POST or None)
    if form.is_valid():
        user = form.save(commit=False)
        username = form.cleaned_data['username']
        password = form.cleaned_data['password']
        user.set_password(password)
        user.save()
        user = authenticate(username=username, password=password)
        if user is not None:
            if user.is_active:
                login(request, user)
                orders = Orders.objects.filter(user_name=request.user.username)
                return render(request, 'event/index.html', {'orders': orders})
    context = {
        "form": form,
    }
    return render(request, 'event/register.html', context)

def track_id(request):
    q = request.POST['trackingnumber']
    error_msg = ''
    orders = Orders.objects.filter(order_id = q)
    return render(request, 'event/order.html', {'orders': orders})

def detail(request, order_id):
    if not request.user.is_authenticated:
        return redirect(reverse('event:index'))
    else:
        user = request.user
        items = Item.objects.filter(order_id=order_id)
        print(items)
        return render(request, 'event/detail.html', {'items': items, 'user': user, 'order_id': order_id})

def change_destination(request, order_id):
    if not request.user.is_authenticated:
        return redirect(reverse('event:index'))
    else:
        form = DestinationForm(request.POST or None, request.FILES or None)
        order = get_object_or_404(Orders, order_id = order_id)
        user_name = order.user_name
        orders = Orders.objects.filter(user_name=request.user.username)
        if form.is_valid():
            if order.o_status == "out for delivery":
                return render(request, 'event/index.html', {'error_message': 'This order is out for delivery, cannot change destination!', 'orders': orders})
            des = form.save(commit=False)
            order.des_x = des.des_x
            order.des_y = des.des_y
            order.save()
            return render(request, 'event/index.html', {'orders': orders})
        context = {
            "form": form,
        }
        return render(request, 'event/change_destination.html', context)