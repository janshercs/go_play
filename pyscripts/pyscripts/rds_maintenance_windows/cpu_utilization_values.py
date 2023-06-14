import boto3
from datetime import datetime, timedelta


session = boto3.Session(profile_name='pd-staging')
staging_cloudwatch = session.client(
    'cloudwatch', region_name='ap-southeast-1')

DAY = datetime(2023, 5, 1)
YESTERDAY = DAY - timedelta(days=1)
HOUR = 60 * 60


def get_lowest_cpu_utilization(data):
    return min(data, key=lambda x: x['Average'])


def get_cpu_utilization_time_percentiles(data) -> list[datetime]:
    data = sorted(data, key=lambda x: x['Average'])
    return [x['Timestamp'] for x in data]


def get_cpu_utilization_time_percentiles_for_cluster(aws_session: boto3.Session, cluster_identifier: str, region: str) -> list[datetime]:
    data = get_cloud_watch_metric(aws_session, cluster_identifier, region)
    return get_cpu_utilization_time_percentiles(data)


def get_cloud_watch_metric(aws_session: boto3.Session, cluster_identifier: str, region: str):
    cloudwatch_client = aws_session.client(
        'cloudwatch', region_name=region)
    response = cloudwatch_client.get_metric_statistics(
        Namespace='AWS/RDS',
        MetricName='CPUUtilization',
        Dimensions=[
            {
                'Name': 'DBClusterIdentifier',
                'Value': cluster_identifier
            },
        ],
        Period=HOUR,
        Statistics=['Average'],
        StartTime=YESTERDAY,
        EndTime=DAY,
    )
    return response['Datapoints']

# sort according to cpu utilization
