import os
from flask import Flask, render_template, request, redirect

app: Flask = Flask(__name__)
# app.config['SECRET_KEY'] =  os.urandom(24)

@app.route('/', methods=['GET'])
def index():

    if request.method == 'GET':
        return render_template('index.html',
            title='API Status Check',
            headline='API Status Check',
            message=''
        )

# Main function is called only when executing ”python app.py”
if __name__ == '__main__':
   app.debug = True
   app.run(host='0.0.0.0', port=80)