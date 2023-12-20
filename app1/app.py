from flask import Flask, render_template
import random

app = Flask(__name__)

@app.route('/')
def generate_random_number():
    # Generate a random number
    random_number = random.randint(1, 100)

    # Generate random background color in hex format
    random_color = "#{:06x}".format(random.randint(0, 0xFFFFFF))

    # Render the HTML template with the random number and background color
    return render_template('index.html', number=random_number, color=random_color)
