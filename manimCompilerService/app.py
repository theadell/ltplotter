from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/post_endpoint', methods=['POST'])
def handle_post():
    data = request.get_json()
    return jsonify({"message": "Data received", "data": data}), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
