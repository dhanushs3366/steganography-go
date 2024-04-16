
import numpy as np
from PIL import Image


import numpy as np
from PIL import Image
def get_lsb_bit(num: int) -> int:
    return num & 1

def decode(image_path: str) -> str:
    img = Image.open(image_path)
    tensor = np.array(img)

    # Flatten the image tensor to process all pixels at once
    flat_tensor = tensor.reshape(-1, 3)

    channels=flat_tensor.shape[-1]
    total_bits=[]
    break_flag=False
    for channel in range(channels):
        pixels=flat_tensor[:,channel]
        lsb_bits=[get_lsb_bit(pixel) for pixel in pixels]
        bit_arr=np.array(lsb_bits).reshape(-1,8)
        null_bit="0b"+"0"*8
        for row in bit_arr:
            bits="0b"
            bit="".join(map(str,row))
            bits+=bit
            if bits==null_bit:
                break_flag=True
                break
            
            total_bits.append(bits)
        if break_flag:
            break
        
    

    total_bits=[int(bit,2) for bit in total_bits]
    
    decode_txt=""
    for bit in total_bits:
        decode_txt+=chr(bit)
    print(decode_txt)
    return decode_txt

# Example usage:
decoded_message = decode("../output/encoded.png")
print("Decoded message:", decoded_message[:50])
