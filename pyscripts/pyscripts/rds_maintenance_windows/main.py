from maintenance_window import get_maintenance_windows
from cpu_utilization_values import get_cpu_utilization_time_percentiles_for_cluster
from rds_statistics import get_formatted_maintenance_window_cpu_percentile_ranking
import boto3

EU_REGION = 'eu-central-1'
AP_REGION = 'ap-southeast-1'
ALL_REGIONS = [EU_REGION, AP_REGION]

session = boto3.Session(profile_name='pd-production')


def main():
    count = 0
    for cluster in get_maintenance_windows(session, EU_REGION):
        while count < 3:
            cpu_utilization_time_percentiles = get_cpu_utilization_time_percentiles_for_cluster(
                session, cluster.cluster_identifier, EU_REGION)

            print(
                f'==============={cluster.cluster_identifier}================')
            for result in get_formatted_maintenance_window_cpu_percentile_ranking(cluster, cpu_utilization_time_percentiles):
                print(result)
            count += 1


if __name__ == "__main__":
    main()
