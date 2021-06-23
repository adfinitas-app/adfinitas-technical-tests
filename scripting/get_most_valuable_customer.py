import sqlite3
from typing import Dict

import pandas


def get_most_valuable_customer() -> Dict[str, int]: return pandas.read_sql_query('SELECT Email, Total FROM customers INNER JOIN invoices ON customers.CustomerId = invoices.CustomerId', sqlite3.connect("../data/chinook.db"))[["Email", "Total"]].groupby("Email").sum().sort_values("Total", ascending=False).head(10).to_dict()["Total"]
