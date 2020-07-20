import os
import subprocess
import ctypes  # An included library with Python install.
from xml.dom import minidom

os.system("PlatformDiscovery.exe > output.xml")
xmldoc = minidom.parse("output.xml")
solution_list = xmldoc.getElementsByTagName("Solution")
for solution in solution_list:
    if solution.attributes["name"].value == "Intel(R) AMT":
        AMTstring = solution.attributes["state"].value
        if AMTstring == "not supported":
            ctypes.windll.user32.MessageBoxW(0, "vPro is not supported on this system", "Information", 1)
        elif AMTstring == "unconfigured":
            ctypes.windll.user32.MessageBoxW(0, "vPro is supported on this system, but not configured", "Information", 1)
        elif AMTstring == "configured":
            ctypes.windll.user32.MessageBoxW(0, "vPro is supported and configured on this system", "Information", 1)
