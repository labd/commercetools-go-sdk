import os

import yaml


class Loader(yaml.SafeLoader):
    def __init__(self, stream):
        try:
            self._root = os.path.split(stream.name)[0]
        except AttributeError:
            self._root = os.path.curdir
        super().__init__(stream)


def construct_include(loader, node):
    filename = os.path.abspath(
        os.path.join(loader._root, loader.construct_scalar(node))
    )
    extension = os.path.splitext(filename)[1].lstrip(".")

    with open(filename, "r") as f:
        if extension in ("yml", "yaml", "raml"):
            return yaml.load(f, Loader)
        return f.read()


yaml.add_constructor("!include", construct_include, Loader)


if __name__ == "__main__":
    import sys

    with open(sys.argv[1], "r") as fh:
        data = yaml.load(fh, Loader)
    print(yaml.dump(data))
