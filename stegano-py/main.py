from flask import Flask, request, jsonify
from encode import encode
from decode import decode


app = Flask(__name__)

@app.route("/encode", methods=["POST"])
def encode_text():
    # if "image_path" not in request.form or "text" not in request.form:
    #     return "Image and text are required", 400
    
    image_path = "../output/encode.png"
    text = request.form["text"]

    encode(image_path, text)
    return jsonify({"message": "Image encoded successfully"}), 200

@app.route("/decode", methods=["POST"])
def decode_image():
    
    image_file = request.form["image_path"]
    
    decoded_text = decode(image_file)
    return decoded_text

@app.route("/",methods=["GET"])
def Home():
    return "<h1>HIII</h1>"
if __name__ == "__main__":
    app.run(debug=True,port=3000)
