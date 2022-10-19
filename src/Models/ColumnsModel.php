<?php

namespace Hightemp\WappSnanimageServer\Models;

use RedBeanPHP\Facade as R;
use Hightemp\WappSnanimageServer\Model;

class ColumnsModel extends Model
{
    public const TABLE = 'columns';

    public const COLUMNS = [
        'created_at' => [],
        'updated_at' => [],
        'name' => [],
        'index_type' => [],
        'type' => [],
        'default_value' => [],
    ];
}