import os
import requests
from flask import Flask, render_template
app = Flask(__name__)
def get_counter(counter_endpoint):
    counter_response = requests.get(counter_endpoint)
    return counter_response.text

def increase_counter(counter_endpoint):
    counter_response = requests.post(counter_endpoint)
    return counter_response.text

@app.route('/')
def hello():
    counter_service = os.environ.get('COUNTER_ENDPOINT', default="http://localhost:6000")
    counter_endpoint = f'{counter_service}/api/counter'
    counter = get_counter(counter_endpoint)

    increase_counter(counter_endpoint)

    return render_template('go.html', value=counter)

