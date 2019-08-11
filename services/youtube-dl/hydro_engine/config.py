import os

def get_save_dir():
    app_data = os.getenv("APPDATA")
    if app_data:
        path = os.path.join(app_data, "ultimate-youtubedl")
        try:
            os.makedirs(path)
            print("Successfully created %s directory" % path)
        except OSError:
            print("Failed to create %s directory" % path)
        else:
            return path
    return ""