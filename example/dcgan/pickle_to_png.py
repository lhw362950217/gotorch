import torch
import torchvision.utils as vutils
import glob, os

files = glob.glob("*.pt")
for file in files:
    directory = os.path.splitext(file)[0]
    if not os.path.exists(directory):
        os.mkdir(directory)

    module = torch.jit.load(file)
    images = list(module.parameters())[0]
    for i in range(10):
        image = images[i].detach().cpu().reshape(28, 28)
        vutils.save_image(image, directory + '/' + str(i) + '.png', normalize=True) 