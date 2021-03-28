import click

@click.command()
@click.option('--name', required=True, type=str)
def cmdline(name):
    click.echo('Hello {name}!'.format(name=name))

if __name__ =='__main__':
    cmdline()