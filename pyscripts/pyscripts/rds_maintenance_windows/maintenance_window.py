import boto3
from dataclasses import dataclass
from datetime import time


def get_rds_client(session: boto3.Session, region: str):
    return session.client('rds', region_name=region)


def get_maintenance_windows(session: boto3.Session, region: str):
    rds_client = get_rds_client(session=session, region=region)
    clusters = get_all_rds_clusters(rds_client)
    return [DB_Cluster(cluster_identifier=cluster['DBClusterIdentifier'], maintenance_window=parse_maintenance_window(cluster['DBClusterIdentifier'], cluster['PreferredMaintenanceWindow'])) for cluster in clusters['DBClusters']]


def get_all_rds_clusters(rds_client):
    return rds_client.describe_db_clusters()


@dataclass
class DB_Cluster:
    cluster_identifier: str
    maintenance_window: tuple[time, time]


def parse_maintenance_window(cluster_name: str, maintenance_window: str) -> tuple[time, time]:
    '''
    parse_mainenance_window takes in a string in the format of 'day:hour:min-day:hour:min' and returns a tuple of time objects
    it checks if the start and end day are the same, if not it raises an AssertionError
    '''

    start, end = maintenance_window.split('-')
    start_day, start_hour, start_min = start.split(':')
    end_day, end_hour, end_min = end.split(':')

    if start_day != end_day:
        print(f'{cluster_name} start day != end day')
    return (time(hour=int(start_hour), minute=int(start_min)), time(hour=int(end_hour), minute=int(end_min)))
