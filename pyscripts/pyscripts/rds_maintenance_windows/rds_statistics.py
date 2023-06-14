from datetime import datetime
from maintenance_window import DB_Cluster
import boto3

EU_REGION = 'eu-central-1'
session = boto3.Session(profile_name='pd-staging')


def get_maintenance_window_cpu_percentile_ranking(cluster: DB_Cluster, cpu_utilization_percentiles: list[datetime]) -> list[int]:
    '''
    get_maintenance_window_cpu_percentile_ranking takes in a DB_cluster object, a datetime list and a percentile and returns the percentile ranking of the cpu utilization of the cluster during the maintenance window
    '''

    start, end = cluster.maintenance_window[0], cluster.maintenance_window[1]

    rank_results = []
    for c, percentile in enumerate(cpu_utilization_percentiles):
        if start < percentile.time() <= end:
            rank_results.append(c)

    return rank_results


def get_formatted_maintenance_window_cpu_percentile_ranking(cluster: DB_Cluster, cpu_utilization_percentiles: list[datetime]) -> list[str]:
    result_string = ["{:.2f}%".format(c / 24 * 100) for c in get_maintenance_window_cpu_percentile_ranking(
        cluster, cpu_utilization_percentiles)]
    return result_string
