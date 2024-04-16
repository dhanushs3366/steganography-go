

import tensorflow as tf
from PIL import Image
import numpy as np

def convert_text_to_bin(text: str) -> str:
    bytes_text = ''.join(format(ord(letter), '08b') for letter in text)
    print(bytes_text)
    return bytes_text

def change_lsb(pixel: np.uint8, bit: str) -> np.uint8:
    lsb_mask = 0b11111110 
    masked_pixel = pixel & lsb_mask
    lsb_pixel = masked_pixel | int(bit)
    return np.uint8(lsb_pixel)

def encode(image_path: str, text: str):
    bytes_text = convert_text_to_bin(text)
    img = Image.open(image_path).convert("RGB")
    tensor = tf.convert_to_tensor(img)
    
    # Flatten the image tensor to process all pixels at once
    flat_tensor = tf.reshape(tensor, (-1, 3))

    # Convert text bits to an array of integers
    text_bits = np.array([int(bit) for bit in bytes_text])
    
    text_bits=np.append(text_bits,[0]*8)
    
    num_pixels_to_modify = min(len(bytes_text), flat_tensor.shape[0])

   
    modified_pixels = flat_tensor.numpy().copy()
    modified_pixels[:num_pixels_to_modify] = np.array([change_lsb(pixel, bit) for pixel, bit in zip(flat_tensor.numpy()[:num_pixels_to_modify], text_bits)])

    # Reshape the modified pixels back to the original image shape
    encoded_img_array = np.reshape(modified_pixels, tensor.shape)

    # Save the encoded image
    encoded_img = Image.fromarray(encoded_img_array.astype(np.uint8))  
    encoded_img.save("../output/encoded.png",format="PNG")
    print("Image saved")


encode("../output/encode.png", "Hello, world!")
