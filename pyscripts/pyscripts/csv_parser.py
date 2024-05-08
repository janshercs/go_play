import csv


TRANSFER_TO_SOURCE = set()
# Opening JSON file
with open('./transfer_to_source.csv') as f:
    for row in csv.reader(f):
        TRANSFER_TO_SOURCE.add(row[0])

PROCESSED = [
    ['reference', 'transfer_to_source', 'status',
        'payment_method_code', 'updated_at']
]

long_file_name = 'payments_with_SAP_invoice_status_2024-05-05_to_2025-05-06_without_invoice_row.csv'
with open(long_file_name, 'r') as f:
    for row in csv.reader(f):
        PROCESSED.append(
            [row[0], str(row[0] in TRANSFER_TO_SOURCE), row[1], row[2], row[3]]
        )

with open('output.csv', 'w', newline='') as f:
    csv_writer = csv.writer(f)
    # Write rows to the CSV file
    for row in PROCESSED:
        csv_writer.writerow(row)
